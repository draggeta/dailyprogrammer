function Out-CardinalClock {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory=$true)]
        [string[]]$TimeStamp
    )
    begin {
        Add-Type -AssemblyName System.Speech
        $speak = New-Object System.Speech.Synthesis.SpeechSynthesizer
        $culture = New-Object System.Globalization.Cultureinfo("en-US")
        # $culture = 'en-US' -as [Globalization.CultureInfo]
    }
    process {
        foreach ($stamp in $TimeStamp) {
            $dateTime = Get-Date $stamp
            #$hour = $dateTime.ToString("%h", $culture)
            #$minute = $dateTime.ToString("mm", $culture)
            $designator = $dateTime.ToString("tt", $culture)
            $hour = ''
            $hour = switch ($dateTime.ToString("%h", $culture)) {
                1 {'one'}
                2 {'two'}
                3 {'three'}
                4 {'four'}
                5 {'five'}
                6 {'six'}
                7 {'seven'}
                8 {'eight'}
                9 {'nine'}
                10 {'ten'}
                11 {'eleven'}
                12 {'twelve'}
            }
            $minute = @()
            $minute = switch ($dateTime.ToString("mm", $culture)) {
                '00'{ ; break }
                { $_.StartsWith('0') } { 'oh'}
                { $_.StartsWith('2') } { 'twenty' }
                { $_.StartsWith('3') } { 'thirty' }
                { $_.StartsWith('4') } { 'forty' }
                { $_.StartsWith('5') } { 'fifty' }
                '10' {'ten'; break}
                '11' {'eleven'; break}
                '12' {'twelve'; break}
                '13' {'thirteen'; break}
                '15' {'fifteen'; break}
                '16' {'sixteen'; break}
                '17' {'seventeen'; break}
                '18' {'eighteen'; break}
                '19' {'nineteen'; break}
                { $_.EndsWith('1') }   { 'one' }
                { $_.EndsWith('2') }   { 'two' }
                { $_.EndsWith('3') }   { 'three' }
                { $_.EndsWith('4') }   { 'four' }
                { $_.EndsWith('5') }   { 'five' }
                { $_.EndsWith('6') }   { 'six' }
                { $_.EndsWith('7') }   { 'seven' }
                { $_.EndsWith('8') }   { 'eight' }
                { $_.EndsWith('9') }   { 'nine' }
            }
            $output = "It's {0} {1}{3}{2}" -f $hour, ($minute -join ' '), $designator, (""," ")[$minute.Count -gt 0]
            $output
            $speak.Speak($output)
        }
    }
}

$test = @(
    '00:00'
    '01:30'
    '12:05'
    '14:01'
    '20:29'
    '21:00'
)

Out-CardinalClock -TimeStamp $test
