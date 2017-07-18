function Get-3Sum () {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory=$true)]
        [int[]]$IntegerArray
    )
    begin {
        $sortArray = $IntegerArray | Sort-Object
        $output = @()
    }
    process {
        for ($i = 0; $i -lt $sortArray.count -2; $i++) {
            $a = $sortArray[$i]
            $start = $i + 1
            $end = $sortArray.count - 1
            while ($start -lt $end) {
                $b = $sortArray[$start]
                $c = $sortArray[$end]
                if ($a +$b + $c -eq 0) {
                    $candidate = $a, $b, $c -join ' '
                    if ($output -notcontains $candidate) {
                        $output += $candidate
                        $candidate
                    }
                    $end = $end - 1
                }
                elseif ($a + $b + $c -gt 0) {
                    $end = $end -1
                }
                else {
                    $start = $start + 1
                }
            }
        }
    }
}

$test1 = @(4, 5, -1, -2, -7, 2, -5, -3, -7, -3, 1)
$test2 = @(-1, -6, -3, -7, 5, -8, 2, -8, 1)
$test3 = @(-5, -1, -4, 2, 9, -9, -6, -1, -7)

Get-3Sum -IntegerArray $test1
Get-3Sum -IntegerArray $test2
Get-3Sum -IntegerArray $test3