# 5.2 The barbershop problem

A barbershop consists of a waiting room with n chairs, and the
barber room containing the barber chair. If there are no customers
to be served, the barber goes to sleep. If a customer enters the
barbershop and all chairs are occupied, then the customer leaves
the shop. If the barber is busy, but chairs are available, then the
customer sits in one of the free chairs. If the barber is asleep, the
customer wakes up the barber. Write a program to coordinate the
barber and the customers.

To make the problem a little more concrete, I added the following information:
* Customer threads should invoke a function named *getHairCut*.
* If a customer thread arrives when the shop is full, it can invoke *balk*,
which does not return.
* The barber thread should invoke *cutHair*.
* When the barber invokes cutHair there should be exactly one thread
invoking *getHairCut* concurrently.