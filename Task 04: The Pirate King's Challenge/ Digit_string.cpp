#include <bits/stdc++.h>
using namespace std;

int countOnesAndThrees(const string& s){
    int cnt = 0;
    for(char c : s){
        if (c == '1' || c == '3') cnt++;
    }
    return cnt;
}

int maxKeepable(const string& s){
    int n = s.size();
    int twosSoFar = 0;
    int onesThreesRemaining = countOnesAndThrees(s);
    int best = twosSoFar + onesThreesRemaining;

    for(int i = 0; i < n; i++){
        char c = s[i];
        if(c == '2'){
            twosSoFar++;
        } else if(c == '1' ||c == '3'){
            onesThreesRemaining--; 
        }
        best = max(best, twosSoFar + onesThreesRemaining);
    }
    return best;
}

void solve(){
    string s;
    cin >> s;
    int n = s.size();
    int kept = maxKeepable(s);
    cout << (n - kept) << "\n";
}

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while (t--) {
        solve();
    }
    return 0;
}