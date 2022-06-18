
(*
  helper
*)
let print_int_arr arr = 
  for i = 0 to Array.length arr - 1 do
    print_int arr.(i);
    print_char ' ';
  done
let int_arr = [|3;6;3;1;3;2|]

(*
  input:  [4; 2; 2; 1]
  output: [1; 2; 4; 6] 
  方針: 入力とは別のarray tmp を用意し、tmpは常にソートされている状態を保つようにする
*)

let insertion_sort arr =
  let l = Array.length arr in
  let tmp = Array.make l 0 in
  (* 0番目の要素だけを先に tmp へいれる *)
  Array.set tmp 0 (Array.get arr 0);
  for i = 1 to l-1 do
    let nw = Array.get arr i in
    let j = ref (i-1) in
    let e = ref true in
    while nw < Array.get tmp !j  && !e do
      Array.set tmp (!j+1) tmp.(!j);
      if !j > 0 
      then j := !j-1
      else e := false
    done;
    if !e == false
    then Array.set tmp (!j) nw
    else Array.set tmp (!j+1) nw;
  done;
  tmp

let conv = insertion_sort int_arr
let () = print_int_arr conv;
