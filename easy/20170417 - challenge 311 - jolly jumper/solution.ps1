function Test-Jolly {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory = $true)]
        [ValidateNotNullOrEmpty()]
        [int[]]$Array
    )
    begin {
        $length = $Array.Length
    }
    process {
        if ($length -eq 1) {
            "$Array JOLLY"
        }
        else {
            [array]$c = @()
            [array]$test = 1..($length - 1)
            for ($i = 0; $i -lt $length - 1; $i++) {
                $c += [math]::Abs($Array[$i] - $Array[$i + 1])
            }
            if ($test.Length -eq $c.Length) {
                $jolly = $true
                foreach ($number in $test) {
                    if ($c -notcontains $number) {
                        $jolly = $false
                    }
                }
                if ($jolly) {
                    "$Array JOLLY"
                }
                elseif (-not $jolly) {
                    "$Array NOT JOLLY"
                }
            }
        }
    }
}


Test-Jolly -Array 1,4,2,3
Test-Jolly -Array 1,4,2,-1,6
Test-Jolly -Array 19,22,24,21
Test-Jolly -Array 19,22,24,25
Test-Jolly -Array 2,-1,0,2
