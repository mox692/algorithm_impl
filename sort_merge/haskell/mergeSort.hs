mergeSort :: [a] -> [a]
mergeSort xs = case xs of
    [] -> []
    _ -> xs

-- test
main = do
    print (mergeSort [1,2,3])
    let l = [1,2,3]
    print (scanList' l (listLen l `div` 2))
    print (scanListAfter' l 0 (listLen l))
    print (splitHalf [1,2,3])
    print (splitHalf [1,2,3,4])
    return ()

--
-- helper
--
splitHalf :: [a] -> ([a], [a])
splitHalf xs = case xs of
    [] -> ([], [])
    h:t ->
        (scanList' xs (len `div` 2), scanListAfter' xs 0 len)
            where len = listLen xs

listLen :: [a] -> Int
listLen xs = case xs of
    [] -> 0
    (h:t) -> 1 + listLen t

-- scan
scanList' :: [a] -> Int -> [a]
scanList' xs n = case xs of
    [] -> []
    (h:t) -> case n of
        0 -> []
        _ -> h:scanList' t (n - 1)

scanListAfter' :: [a] -> Int -> Int -> [a]
scanListAfter' xs n len = case xs of
    [] -> []
    (h:t) -> case () of
        _ 
            | n < center -> scanListAfter' t (n + 1) len
            | n >= center -> h : scanListAfter' t (n + 1) len
            | otherwise -> []
            where center = len `div` 2

