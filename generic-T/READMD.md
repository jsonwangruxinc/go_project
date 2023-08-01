# 什么是泛型
1.泛型即开发过程中编写适用于所有类型中的模版，只有在具体使用的时候才能确定其真正的类型

# 泛型的作用与应用场景
1.增加代码的复用，从同类型的服用到不同类型代码复用。
2.应用于不同类型间的代码复用场景，即不同类型需要写相同的处理逻辑，最适合用泛型。

# 泛型的利弊
1.提高了代码复用率，提高编程效率。
2.不同类型间的代码复用，使得代码风格更加优雅。
3.增加了编译器的负担，降低了编译效率。

# golang泛型怎么使用
1.泛型函数
2.泛型类型
3.泛型接口
4.泛型结构体
5.泛型receiver

# 泛型限制
1.匿名结构体于匿名函数不支持泛型
2.不支持类型断言（if）
3.不支持泛型方法，只能通过receiver来实现方法的泛型处理。
4.~后的类型必须为基本类型，不能为接口类型