$hash = @{
    A = '4'
    B = '6'
    E = '3'
    I = '1'
    L = '|'
    M = '(V)'
    N = '(\)'
    O = '0'
    S = '5'
    T = '7'
    V = '\/'
    W = '`//'
}

function Convert-ToLeet {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory=$true)]
        [string]$String
    )
    begin {
        $originalString = $String
        $String = $String.ToUpper()
    }
    process {
        foreach ($kv in $hash.GetEnumerator()) {
            $String = $String.replace($kv.Key, $kv.Value)
        }
        "$originalString -> $String"
    }
    end {}
}

function Convert-FromLeet {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory=$true)]
        [string]$String
    )
    begin {
        $originalString = $String
        $String = $String.ToUpper()
    }
    process {
        foreach ($kv in $hash.GetEnumerator()) {
            $String = $String.replace($kv.Value, $kv.Key)
        }
        "$originalString -> $String"
    }
    end {}
}
