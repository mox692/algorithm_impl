data Stack' a = Elm a (Stack' a) | Nil
    deriving Show

enqueue' :: Stack' a -> a -> Stack' a
enqueue' xs x = case xs of
    Nil -> Elm x Nil
    Elm e s -> Elm e (enqueue' s x)

push' :: Stack' a -> a -> Stack' a
push' xs e = case xs of
    Nil -> Elm e Nil
    Elm h t -> Elm e (Elm h t)

peek' :: Stack' a -> Option a
peek' xs = case xs of
    Nil -> None
    Elm h t -> Ok h

-- MEMO: 値を2つ返すのをどのように表現するかちょっとわからない.
pop' :: Stack' a -> Stack' a
pop' xs = case xs of
    Nil -> Nil
    Elm h t -> case t of
        Nil -> Nil
        Elm h t -> t

-- test
main = do
    let s = Elm 0 Nil
    print (enqueue' (enqueue' s 1) 2)

    let s = Elm 0 Nil
    print (push' (push' s 1) 2)

    let s = Elm 0 (Elm 1 (Elm 2 Nil))
    print (peek' s)

    return ()


-- helper
data Option a = Ok a | None
    deriving Show

unwrap' :: Option a -> a
unwrap' x = case x of
    None -> error "no value to unwrap"
    Ok e -> e

unwrapOrElse' :: Option a -> a -> a
unwrapOrElse' x defaultVal = case x of
    None -> defaultVal
    Ok e -> e
