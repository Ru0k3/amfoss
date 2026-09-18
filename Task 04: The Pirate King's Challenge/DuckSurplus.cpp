#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t > 0) {
        t--;

        int n;
        cin >> n;

        vector<long long> a(n);
        for (int i = 0; i < n; i++) {
            cin >> a[i];
        }


        vector<long long> stackArr(n);
        int top = 0;

        for (int i = 0; i < n; i++) {
            long long current = a[i];

            while (top > 0 && stackArr[top - 1] > current) {
                current = current + stackArr[top - 1];
                top--;
            }

            stackArr[top] = current;
            top++;
        }

        cout << stackArr[top - 1] << "\n";
    }

    return 0;
}
