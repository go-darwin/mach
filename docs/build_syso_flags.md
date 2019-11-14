From [llvm/llvm-project/compiler-rt/lib/tsan/go/buildgo.sh](https://github.com/llvm/llvm-project/blob/master/compiler-rt/lib/tsan/go/buildgo.sh)

```sh
/Applications/Xcode-beta.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/clang gotsan.cpp -c -o ./race_darwin_amd64.syso -I../rtl -I../.. -I../../sanitizer_common -I../../../include -std=c++11 -Wall -fno-exceptions -fno-rtti -DSANITIZER_GO=1 -DSANITIZER_DEADLOCK_DETECTOR_VERSION=2 -fPIC -Wno-unused-const-variable -Wno-unknown-warning-option -mmacosx-version-min=10.7 -stdlib=libc++ -m64 -DSANITIZER_DEBUG=0 -O3 -fomit-frame-pointer -isysroot/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk -iwithsysroot/usr/include -iwithsysroot../../usr/lib/clang/11.0.0/include
```

```sh
-c -o ./race_darwin_amd64.syso -std=c++11 -Wall -fno-exceptions -fno-rtti -fPIC -Wno-unused-const-variable -Wno-unknown-warning-option -mmacosx-version-min=10.7 -stdlib=libc++ -m64 -O3 -fomit-frame-pointer
```
