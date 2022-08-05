from toolkit.provider.user_agent.generate import OS, Browser, generate_user_agent


def test_generate_user_agent():
    assert OS.Windows.value in generate_user_agent(OS.Windows, platform=None).lower()
    assert OS.Linux.value in generate_user_agent(OS.Linux, platform=None).lower()
    assert OS.Mac.value in generate_user_agent(OS.Mac, platform=None).lower()
    assert OS.IOS.value in generate_user_agent(OS.IOS, platform=None).lower()
    assert OS.Android.value in generate_user_agent(OS.Android, platform=None).lower()

    assert (
        Browser.Safari.value
        in generate_user_agent(browser=Browser.Safari, platform=None).lower()
    )

    assert (
        Browser.Chrome.value
        in generate_user_agent(browser=Browser.Chrome, platform=None).lower()
    )

    assert (
        Browser.Firefox.value
        in generate_user_agent(browser=Browser.Firefox, platform=None).lower()
    )

    assert (
        Browser.Opera.value
        in generate_user_agent(browser=Browser.Opera, platform=None).lower()
    )

    assert (
        Browser.Edge.value
        in generate_user_agent(browser=Browser.Edge, platform=None).lower()
    )
