CMX-Bindings
============

Xamarin bindings for Cisco CMX SDK

Getting started
---------------

### Android
- Download CMX Core Library from [developer.cisco.com](https://developer.cisco.com/site/cmx-mobility-services/downloads/) (requires Cisco dev account)
- Extract and put all .jars into Cisco.Cmx.Droid/Jars
- Open solution in Visual Studio 2013
- Build (requires Xamarin.Android for Visual Studio, and Android API 15 (v4.0.3)
- Copy /bin/Debug/Cisco.Cmx.Droid.dll to your Android app project and reference it

### iOS
- Download CMX Framework from [developer.cisco.com](https://developer.cisco.com/site/cmx-mobility-services/downloads/) (requires Cisco dev account)
- Copy CMX.framework/CMX (file with no extension) to Cisco.Cmx.iOS and rename it to **libCMX.a**
- Open terminal, navigate to `Cisco.Cmx.iOS.Dependencies` and run `make`
- Copy the generated `libCmxDependenciesUniversal.a` to Cisco.Cmx.iOS
- Open `Cisco.Cmx.iOS.csproj` in Xamarin Studio
- Build 
- Copy /bin/Debug/Cisco.Cmx.iOS.dll to your iOS app project and reference it