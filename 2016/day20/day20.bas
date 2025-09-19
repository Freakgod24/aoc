DIM AS _INTEGER64 Blacklist(2048, 2), StartRange, EndRange, Ip, Count

Lines = ReadBlackList(Blacklist(), "input.txt")
QuickSort 0, Lines - 1, Blacklist()

FOR i = 0 TO Lines
  StartRange = Blacklist(i, 0)
  EndRange = Blacklist(i, 1)
  IF StartRange > Ip THEN
    ' PRINT Ip
    ' EXIT FOR
    Count = Count + StartRange - Ip
    Ip = EndRange + 1
  ELSEIF EndRange > Ip THEN
    Ip = EndRange + 1
  END IF
NEXT

PRINT COUNT

FUNCTION ReadBlacklist(Blacklist&&(,), Filename$)
  OPEN Filename$ FOR INPUT AS #1
  WHILE NOT EOF(1)
    INPUT #1, Row$
    p = INSTR(0, Row$, "-")
    Blacklist&&(Lines, 0) = VAL(MID$(Row$, 0, p))
    Blacklist&&(Lines, 1) = VAL(MID$(Row$, p+1, LEN(Row$)))
    Lines = Lines + 1
  WEND
  ReadBlacklist = Lines
END FUNCTION

SUB QuickSort(Start AS LONG, Finish AS LONG, Arr&&())
  Hi = Finish: Lo = Start
  Middle&& = Arr&&((Lo+Hi)/2, 0)
  DO
    DO WHILE Arr&&(Lo, 0) < Middle&&: Lo = Lo + 1: Loop
    DO WHILE Arr&&(Hi, 0) > Middle&&: Hi = Hi - 1: Loop
    IF Lo <= Hi THEN
      SWAP Arr&&(Lo, 0), Arr&&(Hi, 0)
      SWAP Arr&&(Lo, 1), Arr&&(Hi, 1)
      LO = Lo + 1: Hi = Hi - 1
    END IF
  LOOP UNTIL Lo > Hi
  IF Hi > Start THEN CALL QuickSort(Start, Hi, Arr&&())
  IF Lo < Finish THEN CALL QuickSort(Lo, Finish, Arr&&())
END SUB
