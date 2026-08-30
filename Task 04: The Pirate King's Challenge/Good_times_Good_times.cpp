#include <bits/stdc++.h>
using namespace std;

int countDigits(long long x){
    int digits = 0;
    while(x > 0){
        digits++;
        x /= 10;
    }
    return digits;
}

long long buildMultiplier(int len){
    long long y = 1;
    for(int i = 0; i < len; i++) y *= 10;
    y += 1;
    return y;
}

void solve(){
    long long x;
    cin >> x;
    int len = countDigits(x);
    long long y = buildMultiplier(len);
    cout << y << "\n";
}

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while(t--){
        solve();
    }
    return 0;
}