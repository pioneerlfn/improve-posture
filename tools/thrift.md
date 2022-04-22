thrift版本不兼容，项目中用的thrift是0.10，但是当前开发机上，直接用brew安装的thrift版本是0.15或者0.16, 虽然brew可以指定安装thrift的版本，但是只有0.9是可以指定的，其它前后几个版本试了都不行
brew install thrift@0.9

然而，0.9与0.10也是不兼容的，我们仍然需要精确安装0.10版本。那么我的选择只有2个：
拿来主义，从有该版本的人那里拷贝二级制
自力更生，从源码编译安装；

先看拿来主义：
问了同事，同事说他安装成功了，本机上是0.10版本。从同事哪儿拷贝来了二进制，发现并不能正常work, 原因是缺少依赖的库。猜想可能他是动态编译的，如果是静态编译的话应该拷过来就可以直接用。当然我们的OS版本也不相同，静态编译这条路可能是OK的，但不保证100%成功。
简单地拷贝别人的二进制，此路不通(动态编译二进制不行)

自己编译吧。
先尝试按照官方教程安装
https://thrift.apache.org/docs/install/os_x.html
报错，具体错误忘了。。。。。

问了下同事怎么安装的，同事说按下面这个教程，步骤如下：
https://www.jianshu.com/p/2b59f532a492

brew uninstall thrift
cd /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core
git log ./Formula/thrift.rb | less，拿到thrift 0.10的commmitId是174a7d20d60ee094f016facdccefb34961808323
切分支
git checkout 174a7d20d60ee094f016facdccefb34961808323

注释掉35行
vim ./Formula/thrift.rb
#  depends_on "python@2" => :optional

brew install ./Formula/thrift.rb
执行最后一步的时候，报错
undefined method `cellar' for #<BottleSpecification:0x00007fd3d73ec828>
试了网上好多办法，都无法解决这个问题，让人头秃。

既然cellar这个错是brew报的，而brew install也是去执行Formula/thrift.rb, 那么去手动执行这个rb里的操作不就行了么。
thrift.rb长这样：
```ruby
class Thrift < Formula
  desc "Framework for scalable cross-language services development"
  homepage "https://thrift.apache.org/"
  url "https://www.apache.org/dyn/closer.cgi?path=/thrift/0.10.0/thrift-0.10.0.tar.gz"
  sha256 "2289d02de6e8db04cbbabb921aeb62bfe3098c4c83f36eec6c31194301efa10b"


  bottle do
    cellar :any
    rebuild 1
    sha256 "f3bd35df2ba94af77f15a836668db5eb7dfc0d37c77a2f77bc6cc980e1524f27" => :sierra
    sha256 "528061b3a3689341d76d76a7faaa6345100bbbadeb4055a26f1acb6377aad3ba" => :el_capitan
    sha256 "7b1c9edc94356d9cb426237fab09143c64d2bb2a85d86f7d8236702fa110f90c" => :yosemite
  end


  head do
    url "https://git-wip-us.apache.org/repos/asf/thrift.git"


    depends_on "autoconf" => :build
    depends_on "automake" => :build
    depends_on "libtool" => :build
    depends_on "pkg-config" => :build
  end


  option "with-haskell", "Install Haskell binding"
  option "with-erlang", "Install Erlang binding"
  option "with-java", "Install Java binding"
  option "with-perl", "Install Perl binding"
  option "with-php", "Install PHP binding"
  option "with-libevent", "Install nonblocking server libraries"


  depends_on "bison" => :build
  depends_on "boost"
  depends_on "openssl"
  depends_on "libevent" => :optional
  depends_on :python => :optional


  def install
    system "./bootstrap.sh" unless build.stable?


    exclusions = ["--without-ruby", "--disable-tests", "--without-php_extension"]


    exclusions << "--without-python" if build.without? "python"
    exclusions << "--without-haskell" if build.without? "haskell"
    exclusions << "--without-java" if build.without? "java"
    exclusions << "--without-perl" if build.without? "perl"
    exclusions << "--without-php" if build.without? "php"
    exclusions << "--without-erlang" if build.without? "erlang"


    ENV.cxx11 if MacOS.version >= :mavericks && ENV.compiler == :clang


    # Don't install extensions to /usr:
    ENV["PY_PREFIX"] = prefix
    ENV["PHP_PREFIX"] = prefix


    system "./configure", "--disable-debug",
                          "--prefix=#{prefix}",
                          "--libdir=#{lib}",
                          "--with-openssl=#{Formula["openssl"].opt_prefix}",
                          *exclusions
    ENV.deparallelize
    system "make"
    system "make", "install"
  end


  def caveats; <<-EOS.undent
    To install Ruby binding:
      gem install thrift


    To install PHP extension for e.g. PHP 5.5:
      brew install homebrew/php/php55-thrift
  EOS
  end
end
```
ruby语法我不懂，但是看了下，这个rb里的步骤大概分为:

1. 下载源码
2. 执行./bootstrap.sh
2. 执行./configure 生成makefile
3. 执行 make
4. 执行make install 

thirfit历史版本的源码有个存档，可以在这里下载0.10的
http://archive.apache.org/dist/thrift/0.10.0/

make的时候报错了，刚开始报一些bundle版本的错误，根据提示回退到较低版本，还是没法成功。
最后看到configure里有一些without的配置：
```
ruby
    exclusions = ["--without-ruby", "--disable-tests", "--without-php_extension"]

    exclusions << "--without-python" if build.without? "python"
    exclusions << "--without-haskell" if build.without? "haskell"
    exclusions << "--without-java" if build.without? "java"
    exclusions << "--without-perl" if build.without? "perl"
    exclusions << "--without-php" if build.without? "php"
    exclusions << "--without-erlang" if build.wi

```
意思就是可以选择不编译哪些语言，这里我除了保留cpp和Go，其它的全部选择了 --without
则执行configure是这样的：
```shell
./configure --without-ruby  --disable-tests --without-php_extension --without-python  --without-haskell --without-java  --without-perl --without-php --without-erlang --without-nodejs --without-py3

```
再执行make, make install, 没有报错
验证版本：
➜  thrift-0.10.0 thrift --version
```shell
Thrift version 0.10.0
➜  thrift-0.10.0 which thrift
/usr/local/bin/thrift
➜  thrift-0.10.0

```
经历千辛万苦，终于成功了！！！
