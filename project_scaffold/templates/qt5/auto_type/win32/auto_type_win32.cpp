{{SLASH_COMMENTS}}

#include <Windows.h>
#include <array>
#include <functional>
#include <vector>

#include "../auto_type.h"
#include "key_map.h"

constexpr std::array BROWSER_PROCESS_NAMES{"chrome", "firefox", "opera", "browser", "applicationframehost", "iexplore", "edge"};
constexpr std::string_view BROWSER_WINDOW_CLASS = "Chrome_WidgetWin_1";
constexpr SHORT SHORT_MSB = static_cast<SHORT>(0b10000000'00000000);

class AutoType::AutoTypeImpl {
private:
  std::array<bool, MAXBYTE> extended_keys_{};

public:
  AutoTypeImpl() {
    extended_keys_.at(VK_MENU) = true;
    extended_keys_.at(VK_RMENU) = true;
    extended_keys_.at(VK_CONTROL) = true;
    extended_keys_.at(VK_RCONTROL) = true;
    extended_keys_.at(VK_INSERT) = true;
    extended_keys_.at(VK_DELETE) = true;
    extended_keys_.at(VK_HOME) = true;
    extended_keys_.at(VK_END) = true;
    extended_keys_.at(VK_PRIOR) = true;
    extended_keys_.at(VK_NEXT) = true;
    extended_keys_.at(VK_DOWN) = true;
    extended_keys_.at(VK_LEFT) = true;
    extended_keys_.at(VK_RIGHT) = true;
    extended_keys_.at(VK_UP) = true;
    extended_keys_.at(VK_NUMLOCK) = true;
    extended_keys_.at(VK_CANCEL) = true;
    extended_keys_.at(VK_SNAPSHOT) = true;
    extended_keys_.at(VK_DIVIDE) = true;
  }

  bool is_external_key(uint16_t code) {
    return code < extended_keys_.size() && extended_keys_.at(code);
  }

  static HKL active_layout() {
    return GetKeyboardLayout(GetWindowThreadProcessId(GetForegroundWindow(), nullptr));
  }

  static std::optional<KeyCodeWithModifiers> char_to_key_code(HKL layout, char32_t character) {
    if (character > WCHAR_MAX) {
      return std::nullopt;
    }
    if (character == '\n') {
      return KeyCodeWithModifiers{VK_RETURN, Modifier::None};
    }
    auto scan_code_ex = VkKeyScanEx(static_cast<WCHAR>(character), layout);

    auto vk = LOBYTE(scan_code_ex);
    auto shift_state = HIBYTE(scan_code_ex);

    if (!vk || vk == static_cast<BYTE>(-1) || shift_state == static_cast<BYTE>(-1)) {
      return std::nullopt;
    }

    if (shift_state && shift_state != 1) {
      // we properly send only shift now, other modifiers can cause issues
      return std::nullopt;
    }

    KeyCodeWithModifiers res{};
    res.code = vk;

    static constexpr std::array SHIFT_STATES_MODIFIERS{
        std::make_pair(1, Modifier::Shift),
        // can cause issues
        // std::make_pair(2, Modifier::Ctrl),
        // std::make_pair(4, Modifier::Alt),
    };

    for (auto [shift_state_mask, modifier] : SHIFT_STATES_MODIFIERS) {
      if (shift_state & shift_state_mask) {
        res.modifier = res.modifier | modifier;
      }
    }

    return res;
  };
};

AutoType::AutoType() : impl_(std::make_unique<AutoType::AutoTypeImpl>()) {}

AutoType::~AutoType() = default;

AutoTypeResult AutoType::key_move(Direction direction, char32_t character,
                                  std::optional<uint16_t> code, Modifier /*unused*/) {
  auto down = direction == Direction::Down;

  static constexpr auto UTF_HIGH_SURROGATE_START = 0xD800U;
  static constexpr auto UTF_LOW_SURROGATE_START = 0xDC00U;
  static constexpr auto UTF_LOW_SURROGATE_END = 0xDFFFU;
  static constexpr auto UNICODE_PLANE00_END = 0xFFFFU;
  static constexpr auto UNICODE_PLANE01_START = 0x10000UL;
  static constexpr auto UNICODE_PLANE16_END = 0x10FFFFU;
  static constexpr auto DIVIDE_BY = 0x400U;

  std::vector<WORD> chars;
  if (character) {
    if (character > UNICODE_PLANE16_END || (character >= UTF_HIGH_SURROGATE_START && character <= UTF_LOW_SURROGATE_END)) {
      return throw_or_return(AutoTypeResult::BadArg,
                             std::string("Bad character: ") + std::to_string(character));
    }
    if (character <= UNICODE_PLANE00_END) {
      chars.push_back(static_cast<WORD>(character));
    } else {
      chars.push_back(static_cast<WORD>(((character - UNICODE_PLANE01_START) / DIVIDE_BY) + UTF_HIGH_SURROGATE_START));
      chars.push_back(static_cast<WORD>(((character - UNICODE_PLANE01_START) % DIVIDE_BY) + UTF_LOW_SURROGATE_START));
    }
  } else {
    chars.push_back(0);
  }

  std::vector<INPUT> inputs;

  for (auto ch : chars) {
    KEYBDINPUT keyboard_input{};
    keyboard_input.dwFlags = down ? 0 : KEYEVENTF_KEYUP;
    keyboard_input.dwExtraInfo = GetMessageExtraInfo();

    if (code.has_value()) {
      keyboard_input.wVk = code.value();
      if (impl_->is_external_key(code.value())) {
        keyboard_input.dwFlags |= KEYEVENTF_EXTENDEDKEY;
      }
      auto scan_code = MapVirtualKey(code.value(), MAPVK_VK_TO_VSC);
      keyboard_input.wScan = LOWORD(scan_code);
    } else {
      keyboard_input.wVk = 0;
      keyboard_input.wScan = ch;
      keyboard_input.dwFlags |= KEYEVENTF_UNICODE;
    }

    INPUT input = {};
    input.type = INPUT_KEYBOARD;
    input.ki = keyboard_input;

    inputs.push_back(input);
  }

  auto events_sent = SendInput(static_cast<UINT>(inputs.size()), inputs.data(), sizeof(INPUT));
  if (events_sent != inputs.size()) {
    return throw_or_return(AutoTypeResult::OsError, "SendInput error");
  }

  return AutoTypeResult::Ok;
}

Modifier AutoType::get_pressed_modifiers() {
  static constexpr std::array FLAGS_MODIFIERS{
      std::make_pair(VK_LWIN, Modifier::LeftWin),
      std::make_pair(VK_RWIN, Modifier::RightWin),
      std::make_pair(VK_LSHIFT, Modifier::LeftShift),
      std::make_pair(VK_RSHIFT, Modifier::RightShift),
      std::make_pair(VK_LMENU, Modifier::LeftAlt),
      std::make_pair(VK_RMENU, Modifier::RightAlt),
      std::make_pair(VK_LCONTROL, Modifier::LeftCtrl),
      std::make_pair(VK_RCONTROL, Modifier::RightCtrl),
  };
  auto pressed_modifiers = Modifier::None;
  for (auto [key_code, modifier] : FLAGS_MODIFIERS) {
    auto key_state = GetAsyncKeyState(key_code);
    if (key_state & SHORT_MSB) {
      pressed_modifiers = pressed_modifiers | modifier;
    }
  }
  return pressed_modifiers;
}

Modifier AutoType::shortcut_modifier() { return Modifier::Control; }

std::optional<uint16_t> AutoType::os_key_code(KeyCode code) {
  if (code == KeyCode::Undefined) {
    return std::nullopt;
  }
  auto native_key_code = map_key_code(code);
  if (!native_key_code) {
    return std::nullopt;
  }
  return native_key_code;
}

std::optional<KeyCodeWithModifiers> AutoType::os_key_code_for_char(char32_t character) {
  return impl_->char_to_key_code(impl_->active_layout(), character);
}

std::vector<std::optional<KeyCodeWithModifiers>>
AutoType::os_key_codes_for_chars(std::u32string_view text) {
  auto layout = impl_->active_layout();
  std::vector<std::optional<KeyCodeWithModifiers>> result(text.length());
  auto length = text.length();
  for (size_t i = 0; i < length; i++) {
    result[i] = impl_->char_to_key_code(layout, text[i]);
  }
  return result;
}

AutoTypeTextTransaction AutoType::begin_batch_text_entry() { return AutoTypeTextTransaction(); }
