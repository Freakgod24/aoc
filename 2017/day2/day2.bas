Const SHEET_WIDTH = 16
Const FILENAME = "input_day2.txt"

' --------------------
' Main Program
' --------------------

Dim row As Row
Dim values(1 To SHEET_WIDTH) As Integer

Open FILENAME For Input As #1
While Not EOF(1)
    Input #1, strng$
    Row_Read row, strng$, values()
Wend

Print row.checksum1, row.checksum2

' --------------------
' Row Class
' --------------------

Type Row
    min_value As Integer
    max_value As Integer
    checksum1 As Long
    checksum2 As Long
    value_index As Integer
    split_index As Integer
End Type

Sub Row_Reset (self As Row)
    self.min_value = _INTEGER_MAX&
    self.max_value = 0
    self.value_index = 1
    self.split_index = 1
    ' We preserve the checksums between rows
End Sub

Sub Row_Read (self As Row, strng As String, values() As Integer)

    Row_Reset self

    For i = 1 To Len(strng)
        c$ = Mid$(strng, i, 1)
        If Asc(c$) = _ASC_HT Or i = Len(strng) Then
            ' Last item doesn't have a \t to exclude at the end
            If i = Len(strng) Then
                value = ValueFromString(strng, self.split_index, i - self.split_index + 1)
            Else
                value = ValueFromString(strng, self.split_index, i - self.split_index)
            End If

            Row_UpdateValue self, value, values()
            Row_UpdateMinMax self, value

            self.split_index = i + 1
        End If
    Next

    Row_UpdateChecksum1 self
    Row_UpdateChecksum2 self, values()
End Sub

Sub Row_UpdateValue (self As Row, value As Integer, values() As Integer)
    values(self.value_index) = value
    self.value_index = self.value_index + 1
End Sub

Sub Row_UpdateMinMax (self As Row, value As Integer)
    self.min_value = Min(self.min_value, value)
    self.max_value = Max(self.max_value, value)
End Sub

Sub Row_UpdateChecksum1 (self As Row)
    self.checksum1 = self.checksum1 + self.max_value - self.min_value
End Sub

Sub Row_UpdateChecksum2 (self As Row, values() As Integer)
    For i = 1 To SHEET_WIDTH - 1
        For j = i + 1 To SHEET_WIDTH
            If (values(i) Mod values(j)) = 0 Then
                self.checksum2 = self.checksum2 + values(i) / values(j)
                Exit Sub
            ElseIf (values(j) Mod values(i)) = 0 Then
                self.checksum2 = self.checksum2 + values(j) / values(i)
                Exit Sub
            End If
        Next
    Next
End Sub

' --------------------
' Utilities
' --------------------

Function ValueFromString (strng As String, start_index As Integer, stop_index As Integer)
    ValueFromString = Val(Mid$(strng, start_index, stop_index))
End Function

Function Min (a As Integer, b As Integer)
    If a < b Then Min = a Else Min = b
End Function

Function Max (a As Integer, b As Integer)
    If a > b Then Max = a Else Max = b
End Function
