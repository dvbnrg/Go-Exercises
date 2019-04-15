
Given a list of tuples T = [(k1 -> -v), (k2 -> +v), ...], where **k** is the identification key of users
in a payment system and **v** is the value of debit or credit transactions in this system.

Considering that the transactions need to be sent out to an external system B, and each user
has several transactions on the list T, and that B only accepts messages that represent a
transfer of value from a payer user to a payee user. 

Write a function that reduces the
transactions on list T to a new list T' that contain the messages to be sent to the external system

B. Assume that all transactions in T will sum up to 0. Be aware that B charges a fee for every
element of T'.

Sample input:
T = [(u1 -> -10), (u2 -> +10), (u2 -> -3), (u3 -> +3), (u2 -> -1), (u1 -> +1)]

Sample output:
T' = [(u1, u2, 9), (u2, u3, 3)]

Implement your solution using one of the following programming languages: Scala, C, C++,
Rust, or Go.

Add a comment with the computational complexity (Big-O) of your solution.