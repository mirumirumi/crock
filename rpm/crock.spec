%define _binaries_in_noarch_packages_terminate_build 0

Summary: crock is rock clock
Name:    crock
Version: #####
Release: 1
License: MIT
Group:   Applications/System
URL:     https://github.com/mirumirumi/crock

Source0:   %{name}
Source1:   %{name}.initd
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root

%description
%{summary}

%prep

%build

%install
%{__rm} -rf %{buildroot}
%{__install} -Dp -m0755 %{SOURCE0} %{buildroot}/usr/local/bin/%{name}
%{__install} -Dp -m0755 %{SOURCE1} %{buildroot}/%{_initrddir}/%{name}

%clean

%post
/sbin/chkconfig --add %{name}

%files
%defattr(-,root,root)
/usr/local/bin/%{name}
%{_initrddir}/%{name}
