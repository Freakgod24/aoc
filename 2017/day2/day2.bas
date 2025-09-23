Dim As Long checksum, row_min, row_max

Open "input_day2.txt" For Input As #1
While Not EOF(1)
    Input #1, Row$

    s = 0
    row_min = 2147483647
    row_max = 0


    For i = 1 To Len(Row$)
        c$ = Mid$(Row$, i, 1)
        If Asc(c$) = 9 Then
            n = Val(Mid$(Row$, s, i - s))
            row_min = Min(row_min, n)
            row_max = Max(row_max, n)
            s = i + 1
        End If
        If i = Len(Row$) Then
            n = Val(Mid$(Row$, s, i - s + 1))
            row_min = Min(row_min, n)
            row_max = Max(row_max, n)
        End If
    Next
    checksum = checksum + row_max - row_min
    Print row_min, row_max, checksum
Wend

Function Min (a As Long, b As Long)
    If a < b Then Min = a Else Min = b
End Function

Function Max (a As Long, b As Long)
    If a > b Then Max = a Else Max = b
End Function
