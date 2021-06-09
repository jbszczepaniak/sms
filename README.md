### SMS - Socks Management System

SMS helps to manage your socks in an efficent way. It's united tested and I am running it in production for a while without any outages. I invite you to learn about SMS.

Basically this helps you to go from this:

![before](./before.png)

to this:

![after](./after.png)

or if you really like perfection, to this (it requires ironing socks, it's for pro users only):

![after_extreme](./after_extreme.png)

## Idea

It is very simple. The idea assumes that you spend time on searching for socks anyway.

> Instead of matching socks together when you're in a hurry, match them in the peace.

## Setup

You need two organizers. IKEA SKUBB organizers are alright, you can use them.

## The algorithm

Organizer #1 is going to be a queue. Organizer #2 is going to be just a box of unordered socks. Whenever you do the loundry, you take all your fresh socks and put it in Organizer #2. Then you take all of the socks from Organizer #2 one by one, and whenever you find a matching pair you enqueue Organizer #1 with that pair. After finishing, if you still have socks without a pair, you put them back in the Organizer#2. In the morning, you dequeue Organizer#1. If there is nothing in the queue - it's clear sign you need to do the laundry.

## Why not just roll the socks into pairs and put them into box?

Because it ruins your socks. The elastic band that ends the sock loosens and it gets ugly, you don't want that. Also when you are choosing your socks in the mornign randomly from the box you wear some of the socks more frequently than others - it's bad. Wearing off all of the socks in the same pace is much better.

## Limitations

1. After loosing a sock, you will never get the other one from this system. We would need to have some way to keep track of socks that never go from Organizer#2 to Organizer#1.