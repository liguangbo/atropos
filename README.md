#Atropos
Atropos is a DSL for edit JSON struct.

# Design

Antlr for grammar

We need a DSL can edit JSON result according to the http request.

Using in client strategy deliver basically.

## Assignment

```
a = 1
b = ['test', 2, false]
c = [b, ['test2', 4, true]]
c[2] = 3
c[0][1] = 3

// stuct must defiened outside
input.a = 1
```

## Expression

```
a == 4 
a != 4
a > 4
a >= 4
a < 4
a <= 4
a regexpMatch '\d'
a = a + b
a = a - b
a = a * b
a = a / b
a = a % b
a = !a
a = -a
a = a * (a + a)
a contains b // is set a contains b, b could be an array or an object
a !contains b // is set a not contains b, b could be an array or an object
a && b
a || b
a && (b || c)
```

## If statement

```
if expression then
// do something
end

if expression then
// do something
elseif expression then
// do something
end

if expression then
// do something
else
// do something
end
```

## Return
return anywhere
```
return
```

## Function
a.count()
a.avg()
a.max()
a.min()
a.sum()
# More

read and run in example_test.go for more examples