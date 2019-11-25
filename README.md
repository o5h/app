# Simple framevork for Golang/Egl/OpenglES2.0 application

Depends on

```text
d3dcompiler_47.dll
libEGL.dll
libGLESv2.dll
```

Dlls can be found in [Chromium Embedded Framework](http://opensource.spotify.com/cefbuilds/cef_binary_78.3.4%2Bge17bba6%2Bchromium-78.0.3904.108_windows64_minimal.tar.bz2)


## mingw64 needed
- Download and run the installer from http://www.msys2.org/
- In the mingw64 msys2 console, run the following:
  - `pacman -Syu`
  - `pacman -Su`
  - `pacman -S mingw64/mingw-w64-x86_64-gcc mingw64/mingw-w64-x86_64-go mingw64/mingw-w64-x86_64-pkg-config msys/git`
   
* clang 
   `pacman -S mingw-w64-x86_64-clang`
