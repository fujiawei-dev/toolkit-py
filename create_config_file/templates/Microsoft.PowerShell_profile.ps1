<#
 * Author: Rustle Karl
 * Email: fu.jiawei@outlook.com
 * Date: 2020.08.07
 * Copyright: No copyright. You can use this code for anything with no warranty.
#>

#------------------------------- Import Modules BEGIN -------------------------------

# 引入 posh-git
Import-Module posh-git

# 引入 oh-my-posh
Import-Module oh-my-posh

# 设置 PowerShell 主题
Set-Theme robbyrussell

#------------------------------- Import Modules END   -------------------------------


#-------------------------------  Set Hot-keys BEGIN  -------------------------------

# 设置 Tab 键补全
Set-PSReadlineKeyHandler -Key "Tab" -Function Complete
Set-PSReadLineKeyHandler -Key "Tab" -Function MenuComplete

# 设置 Ctrl+d 为退出 PowerShell
Set-PSReadlineKeyHandler -Key "Ctrl+d" -Function ViExit

# 设置 Ctrl+z 为撤销
Set-PSReadLineKeyHandler -Key "Ctrl+z" -Function Undo

# 设置 Ctrl+f 为确认当前提示
Set-PSReadLineKeyHandler -Key "Ctrl+f" -Function ForwardWord

# 设置向上键为后向搜索历史记录
Set-PSReadLineKeyHandler -Key UpArrow -Function HistorySearchBackward

# 设置向下键为前向搜索历史纪录
Set-PSReadLineKeyHandler -Key DownArrow -Function HistorySearchForward

# 类似 Fish Shell 的输入提示
Set-PSReadLineOption -PredictionSource History

#-------------------------------  Set Hot-keys END    -------------------------------


#-------------------------------   Set Alias BEGIN     -------------------------------

# 目录列表 ls & ll
function ListDirectory {
    (Get-ChildItem).Name
    Write-Host("")
}
Set-Alias -Name ls -Value ListDirectory
Set-Alias -Name ll -Value Get-ChildItem

# pwd 显示当前目录
Set-Alias -Name pwd -Value Get-Location

# ffmpeg 默认覆盖文件，不输出标语
function MFFmpeg { ffmpeg $args -hide_banner -y }
Set-Alias -Name ffmpeg -Value MFFmpeg

function MFFprobe { ffprobe -hide_banner $args }
Set-Alias -Name ffprobe -Value MFFprobe

# py 简写 python
Set-Alias -Name py -Value python

# get 简写 git clone
function GitClone { git clone $args }
Set-Alias -Name get -Value GitClone

# gad 简写 git add
function GitAdd { git add $args }
Set-Alias -Name gad -Value GitAdd

# gcmt 简写 git commit -m "args"
function GitCommit { git commit -m $args }
Set-Alias -Name gcmt -Value GitCommit

# gst 简写 git status
function GitStatus { git status }
Set-Alias -Name gst -Value GitStatus

# gph 简写 git push
function GitPush { git push }
Set-Alias -Name gph -Value GitPush

# gh 简写 git hist
function GitHistory { git hist }
Set-Alias -Name gh -Value GitHistory

# md5
function MD5File {
    certutil -hashfile $args MD5
}
Set-Alias -Name md5 -Value MD5File

# rs 启动录音服务
function RecorderServer { D:\OneDrive\Repositories\projects\edge-tts\RecorderClient\RecorderServer.exe }
Set-Alias -Name rs -Value RecorderServer

# viewer 启动浏览器
function StartViewer { python D:\OneDrive\Repositories\projects\auto-sign-in\main.py viewer }
Set-Alias -Name viewer -Value StartViewer

Set-Alias -Name watcher -Value "D:\OneDrive\Market\Program\daily-stocks\watch.bat"
Set-Alias -Name monitor -Value "D:\OneDrive\Market\Program\daily-stocks\monitor.bat"

#-------------------------------    Set Alias END     -------------------------------
