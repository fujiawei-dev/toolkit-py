from user_agent.generate import Browser, OS, generate_user_agent


def test_generate_user_agent():
    print(generate_user_agent())
    print(generate_user_agent(OS.Windows, platform=None))
    print(generate_user_agent(OS.Linux, platform=None))
    print(generate_user_agent(OS.Mac, platform=None))
    print(generate_user_agent(OS.IOS, platform=None))
    print(generate_user_agent(OS.Android, platform=None))
    print(generate_user_agent(browser=Browser.Safari, platform=None))
    print(generate_user_agent(browser=Browser.Chrome, platform=None))
    print(generate_user_agent(browser=Browser.Firefox, platform=None))
    print(generate_user_agent(browser=Browser.Opera, platform=None))
    print(generate_user_agent(browser=Browser.Edge, platform=None))
