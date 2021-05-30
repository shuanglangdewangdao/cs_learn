# 简单工厂模式
主要特点是需要在工厂类中做判断，从而创造相应的产品，当增加新产品时，需要修改工厂类。使用简单工厂模式，我们只需要知道具体的产品型号就可以创建一个产品。

多个产品类

工厂来判断生产什么产品

缺点：工厂类集中了所有产品类的创建逻辑，如果产品量较大，会使得工厂类变的非常臃肿。

    /*
    关键代码：创建过程在工厂类中完成。
    */

    #include <iostream>​
    using namespace std;

    //定义产品类型信息
    typedef enum
    {
        Tank_Type_56,
        Tank_Type_96,
        Tank_Type_Num
    }Tank_Type;
    ​
    //抽象产品类
    class Tank
    {
    public:
        virtual const string& type() = 0;
    };

    //具体的产品类
    class Tank56 : public Tank
    {
    public:
        Tank56():Tank(),m_strType("Tank56")
        {
        }

        const string& type() override
        {
            cout << m_strType.data() << endl;
            return m_strType;
        }
    private:
        string m_strType;
    };

    //具体的产品类
    class Tank96 : public Tank
    {
    public:
        Tank96():Tank(),m_strType("Tank96")
        {
        }
        const string& type() override
        {
            cout << m_strType.data() << endl;
            return m_strType;
        }

    private:
        string m_strType;
    };

    //工厂类
    class TankFactory
    {
    public:
        //根据产品信息创建具体的产品类实例，返回一个抽象产品类
        Tank* createTank(Tank_Type type)
        {
            switch(type)
            {
            case Tank_Type_56:
                return new Tank56();
            case Tank_Type_96:
                return new Tank96();
            default:
                return nullptr;
            }
        }
    };
​
    int main()
    {
        TankFactory* factory = new TankFactory();
        Tank* tank56 = factory->createTank(Tank_Type_56);
        tank56->type();
        Tank* tank96 = factory->createTank(Tank_Type_96);
        tank96->type();

        delete tank96;
        tank96 = nullptr;
        delete tank56;
        tank56 = nullptr;
        delete factory;
        factory = nullptr;

        return 0;
    }
## 工厂模式

定义一个创建对象的接口，其子类去具体现实这个接口以完成具体的创建工作。如果需要增加新的产品类，只需要扩展一个相应的工厂类即可。

将产品和工厂的特征抽象出来，作为接口，由具体的产品类和工厂类来实现

缺点：产品类数据较多时，需要实现大量的工厂类，这无疑增加了代码量。

    /*
    关键代码：创建过程在其子类执行。
    */

    #include <iostream>​
    using namespace std;

    //产品抽象类
    class Tank
    {
    public:
        virtual const string& type() = 0;
    };​
    //具体的产品类
    class Tank56 : public Tank
    {
    public:
        Tank56():Tank(),m_strType("Tank56")
        {
        }

        const string& type() override
        {
            cout << m_strType.data() << endl;
            return m_strType;
        }
    private:
        string m_strType;
    };

    //具体的产品类
    class Tank96 : public Tank
    {
    public:
        Tank96():Tank(),m_strType("Tank96")
        {
        }
        const string& type() override
        {
            cout << m_strType.data() << endl;
            return m_strType;
        }

    private:
        string m_strType;
    };
​
    //抽象工厂类，提供一个创建接口
    class TankFactory
    {
    public:
        //提供创建产品实例的接口，返回抽象产品类
        virtual Tank* createTank() = 0;
    };

    //具体的创建工厂类，使用抽象工厂类提供的接口，去创建具体的产品实例
    class Tank56Factory : public TankFactory
    {
    public:
        Tank* createTank() override
        {
            return new Tank56();
        }
    };

    //具体的创建工厂类，使用抽象工厂类提供的接口，去创建具体的产品实例
    class Tank96Factory : public TankFactory
    {
    public:
        Tank* createTank() override
        {
            return new Tank96();
        }
    };

    int main()
    {
        TankFactory* factory56 = new Tank56Factory();
        Tank* tank56 = factory56->createTank();
        tank56->type();
        TankFactory* factory96 = new Tank96Factory();
        Tank* tank96 = factory96->createTank();
        tank96->type();
        delete tank96;
        tank96 = nullptr;
        delete factory96;
        factory96 = nullptr;

        delete tank56;
        tank56 = nullptr;
        delete factory56;
        factory56 = nullptr;

        return 0;
    }

## 抽象工厂模式
抽象工厂模式提供创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。

当存在多个产品系列，而客户端只使用一个系列的产品时，可以考虑使用抽象工厂模式。

缺点：当增加一个新系列的产品时，不仅需要现实具体的产品类，还需要增加一个新的创建接口，扩展相对困难。

将系列的特征抽象出来作为接口，再进行实现

## 策略模式
策略模式是指定义一系列的算法，把它们单独封装起来，并且使它们可以互相替换，使得算法可以独立于使用它的客户端而变化，也是说这些算法所完成的功能类型是一样的，对外接口也是一样的，只是不同的策略为引起环境角色环境角色表现出不同的行为。

相比于使用大量的if...else，使用策略模式可以降低复杂度，使得代码更容易维护。

缺点：可能需要定义大量的策略类，并且这些策略类都要提供给客户端。

\[环境角色\]  持有一个策略类的引用，最终给客户端调用。

抽象策略作为接口，再进行具体实现，让环境角色持有相应的策列类的引用

## 适配器模式
适配器模式可以将一个类的接口转换成客户端希望的另一个接口，使得原来由于接口不兼容而不能在一起工作的那些类可以在一起工作。通俗的讲就是当我们已经有了一些类，而这些类不能满足新的需求，此时就可以考虑是否能将现有的类适配成可以满足新需求的类。适配器类需要继承或依赖已有的类，实现想要的目标接口。

缺点：过多地使用适配器，会让系统非常零乱，不易整体进行把握。比如，明明看到调用的是 A 接口，其实内部被适配成了 B 接口的实现，一个系统如果太多出现这种情况，无异于一场灾难。因此如果不是很有必要，可以不使用适配器，而是直接对系统进行重构。