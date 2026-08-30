#include <bits/stdc++.h>
using namespace std;

vector<long long> generatePrimes(int count){
    int limit = 130000; 
    vector<bool> isComposite(limit + 1, false);
    vector<long long> primes;

    for(int i = 2; i <= limit; i++){
        if(!isComposite[i]){
            primes.push_back(i);
            if((int)primes.size() == count) break;

            for(long long j = (long long)i * i; j <= limit; j += i){
                isComposite[j] = true;
            }
        }
    }
    return primes;
}

vector<long long> buildSequence(int n, vector<long long>& primes){
    vector<long long> a(n + 1); 
    a[1] = primes[0];
    for(int i = 2; i <= n - 1; i++){
        a[i] = primes[i - 2] * primes[i - 1];
    }
    if(n >= 2){
        a[n] = primes[n - 2];
    }
    return a;
}

void solve(vector<long long>& primes){
    int n;
    cin >> n;

    vector<long long> a = buildSequence(n, primes);

    for(int i = 1; i <= n; i++){
        cout << a[i] << " \n"[i == n];
    }
}

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    vector<long long> primes = generatePrimes(10000);
    int t;
    cin >> t;
    while(t--){
        solve(primes);
    }
    return 0;
}