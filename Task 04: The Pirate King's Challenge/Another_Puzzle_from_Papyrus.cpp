#include <bits/stdc++.h>
using namespace std;

long long costWithoutReorder(vector<int>& a, vector<int>& b){
    int n = a.size();
    long long total = 0;

    for(int i = 0; i < n; i++){
        if (a[i] < b[i]) return -1; 
        total += (a[i] - b[i]);
    }
    return total;
}

long long costWithReorder(vector<int> a, vector<int> b, int c){
    int n = a.size();
    sort(a.begin(), a.end());
    sort(b.begin(), b.end());
    long long total = c; 
    for(int i = 0; i < n; i++){
        if (a[i] < b[i]) return -1;
        total += (a[i] - b[i]);
    }
    return total;
}

void solve(){
    int n, c;
    cin >> n >> c;
    vector<int> a(n), b(n);
    for (int i = 0; i < n; i++) cin >> a[i];
    for (int i = 0; i < n; i++) cin >> b[i];
    long long option1 = costWithoutReorder(a, b);
    long long option2 = costWithReorder(a, b, c);
    long long best = -1;
    if(option1 != -1) best = option1;
    if (option2 != -1){
        if(best == -1 || option2 < best) best = option2;
    }
    cout << best << "\n";
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