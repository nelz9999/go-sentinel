# go-sentinel
Go code experiments for providing constant sentinel error values

# Inspiration
As someone who isn't super well-informed about _programming language design_, I spend a lot of time reading the output of other people who are, and specifically I find [Dave Cheney's writing](https://dave.cheney.net/) to be super approachable and valuable.

In addition to having the privileged of a lot of time to read blogs, I also have time to write small code projects, and after reading Dave's blog post about [constant errors](https://dave.cheney.net/2016/04/07/constant-errors), I thought I'd knock out the reference code, even if this pattern was declared to be [not the best practice](https://dave.cheney.net/2014/12/24/inspecting-errors).

I also tried to see if I could come up with a way to address [a concern with this pattern](https://github.com/golang/go/issues/17226#issuecomment-309125918) that was raised on the (now closed) [proposal to add this pattern to the standard library](https://github.com/golang/go/issues/17226). (It turns out the string equality concern does not seem to be _structurally_ fixable within Go, and falls back to relying on best practices and conventions for descriptive error messages.)
