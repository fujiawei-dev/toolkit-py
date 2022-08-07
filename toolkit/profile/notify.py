def ringtone_reminder(sound="completed.wav"):
    """铃声通知提醒"""
    import winsound

    winsound.PlaySound(sound, winsound.SND_FILENAME)
