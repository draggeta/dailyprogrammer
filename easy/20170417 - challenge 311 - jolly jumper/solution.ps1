function Test-Jolly {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory = $true)]
        [ValidateNotNullOrEmpty()]
        [int[]]$Array
    )
    begin {
        $l = $Array.Length
    }
    process {
        if ($l -eq 1) {
            "$Array JOLLY"
        }
        else {
            [array]$c = @()
            [array]$t = 1..($l - 1)
            for ($i = 0; $i -lt $l - 1; $i++) {
                $c += [math]::Abs($Array[$i] - $Array[$i + 1])
            }
            if (
                ($t | Where-Object {$c -notcontains $_}) -eq $null -and
                ($c | Where-Object {$t -notcontains $_}) -eq $null
            ) {
                "$Array JOLLY"
            }
            else {
                "$Array NOT JOLLY"
            }
        }
    }
}


Test-Jolly -Array 1,4,2,3
Test-Jolly -Array 1,4,2,-1,6
Test-Jolly -Array 19,22,24,21
Test-Jolly -Array 19,22,24,25
Test-Jolly -Array 2,-1,0,2
