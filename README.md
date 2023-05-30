# GoElapseChain - A Blockchain based on ProofOfElapsedTime in Golang

Hey there, folks! 👋 This here's GoElapseChain, a blockchain written in Go. And guess who made it? Yep, yours truly, Mo Ashouri (ashourics@gmail.com). So let's dive in and see what's under the hood!

## What's a Blockchain?

A blockchain, in simple words, is like a chain of digital "blocks", each carrying some data. The really cool thing is that each block's got the hash (like a digital fingerprint) of the previous block. This creates a chain that is super tough to mess with, 'cause if you change something in a block, the hash changes and everyone knows somethin's fishy!

## Architecture

This is a pretty simple setup. We've got the `GoElapseChainBlock` structure that represents each block in the chain. It's got the index (like its position in the chain), timestamp (when it was created), data (the payload), hash (its own fingerprint), previous hash (the fingerprint of the previous block in the chain) and a signature.

The entire blockchain is a slice of these `GoElapseChainBlock` blocks. Super neat, huh?

We also got a private key that's used to sign each block, adding an extra layer of security to the mix.

## Cryptography

Ah, the secret sauce! The security of our blockchain rests on some nifty crypto tricks. First off, we got SHA-256 hashes. This is what gives each block (and the data inside it) a unique fingerprint. If even a tiny bit of the data changes, the hash changes completely.

We're also signing each block using a private key in a process called "signing". This ensures that the block hasn't been tampered with after it was created. If someone messes with the data in the block, the signature won't match and we'll know somethin's up!

## Consensus Algorithm

Our consensus algorithm is based on Proof of Elapsed Time (PoET). It's a simple version, we just wait a random amount of time and then return the time waited. It's not super sophisticated, but hey, it does the job for this little project!

## Wanna Learn More About Blockchain?

Great news, folks! I'm offering a comprehensive course on blockchain design. If you've enjoyed exploring GoElapseChain and wanna dive deeper into the world of blockchains, this course is just for you!

It doesn't matter if you're an experienced developer looking to expand your skills, or a curious newbie who's just getting started with coding. This course will take you on a fun and exciting journey through the world of blockchain technology!

Interested? Just drop me an email at ashourics@gmail.com and I'll get back to you with the deets. 

## Wrapping Up

So there you have it, a whistle-stop tour of GoElapseChain. It's got all the essentials of a blockchain, wrapped up in some neat Go code. Hope you find it helpful, and feel free to reach out if you've got any questions, suggestions or just wanna chat about blockchains!

-- Mo Ashouri (ashourics@gmail.com)
