#include <bits/stdc++.h>
using namespace std;

long long computeLargestPile(vector<long long>& a){
    vector<long long> stack;
    for(long long x : a){
        long long current = x;
        while(!stack.empty() && stack.back() > current){
            current += stack.back();
            stack.pop_back();
        }
        stack.push_back(current);
    }
    return stack.back();
}

void solve(){
    int n;
    cin >> n;
    vector<long long> a(n);
    for(int i = 0; i < n; i++) cin >> a[i];
    cout << computeLargestPile(a) << "\n";
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