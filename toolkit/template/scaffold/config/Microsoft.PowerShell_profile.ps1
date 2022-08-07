<#
 * Author: Rustle Karl
 * Email: fu.jiawei@outlook.com
 * Date: 2020.08.07
 * Copyright: No copyright. You can use this code for anything with no warranty.
#>

#------------------------------- Set Environment Variables BEGIN -------------------------------

$Env:POSHGIT_CYGWIN_WARNING = 'off'

$Env:HTTP_PROXY = 'http://127.0.0.1:7890'
$Env:HTTPS_PROXY = 'http://127.0.0.1:7890'

#------------------------------- Set Environment Variables END -------------------------------

#------------------------------- Import Modules BEGIN -------------------------------

Import-Module posh-git

oh-my-posh init pwsh --config ~/.config/oh-my-posh/themes/robbyrussel.omp.json | Invoke-Expression

#------------------------------- Import Modules END   -------------------------------


#-------------------------------  Set Hot-keys BEGIN  -------------------------------

Set-PSReadlineKeyHandler -Key "tab" -Function Complete
Set-PSReadLineKeyHandler -Key "tab" -Function MenuComplete

Set-PSReadlineKeyHandler -Key "ctrl+d" -Function ViExit
Set-PSReadLineKeyHandler -Key "ctrl+z" -Function Undo
Set-PSReadLineKeyHandler -Key "ctrl+f" -Function ForwardWord
Set-PSReadLineKeyHandler -Key UpArrow -Function HistorySearchBackward
Set-PSReadLineKeyHandler -Key DownArrow -Function HistorySearchForward

Set-PSReadLineOption -PredictionSource History

#-------------------------------  Set Hot-keys END    -------------------------------


#-------------------------------   Set Alias BEGIN     -------------------------------

function ListDirectory
{
    (Get-ChildItem).Name
    Write-Host("")
}
Set-Alias -Name ls -Value ListDirectory

Set-Alias -Name ll -Value Get-ChildItem
Set-Alias -Name pwd -Value Get-Location

function GitClone
{
    git clone $args
}
Set-Alias -Name get -Value GitClone

function GitAdd
{
    git add $args
}
Set-Alias -Name gad -Value GitAdd

function GitCommit
{
    git commit -m $args
}
Set-Alias -Name gcmt -Value GitCommit

function GitStatus
{
    git status
}
Set-Alias -Name gst -Value GitStatus

function GitPush
{
    git push
}
Set-Alias -Name gph -Value GitPush

function GitPull
{
    git pull $args
}
Set-Alias -Name gpl -Value GitPull

function GitHistory
{
    git hist
}
Set-Alias -Name gh -Value GitHistory

function MD5File
{
    certutil -hashfile $args MD5
}
Set-Alias -Name md5 -Value MD5File

function ChocoInstall
{
    choco install $args
}
Set-Alias -Name apti -Value ChocoInstall

function ChocoUninstall
{
    choco uninstall $args
}
Set-Alias -Name aptr -Value ChocoUninstall

function ChocoSearch
{
    choco search $args
}
Set-Alias -Name apts -Value ChocoSearch

function ChocoUpgradeAll
{
    choco upgrade all $args
}
Set-Alias -Name aptu -Value ChocoUpgradeAll

#-------------------------------    Set Alias END     -------------------------------
