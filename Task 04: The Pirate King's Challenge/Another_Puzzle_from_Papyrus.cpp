#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t > 0) {
        t--;

        int n, c;
        cin >> n >> c;

        vector<int> a(n);
        vector<int> b(n);

        for (int i = 0; i < n; i++) cin >> a[i];
        for (int i = 0; i < n; i++) cin >> b[i];

        // no reorder, match positions as given.
        long long option1 = 0;
        bool option1Possible = true;

        for (int i = 0; i < n; i++) {
            if (a[i] < b[i]) {
                option1Possible = false;
            } else {
                option1 = option1 + (a[i] - b[i]);
            }
        }

        // pay c, then sort both and match smallest-to-smallest.
        vector<int> aSorted = a;
        vector<int> bSorted = b;
        sort(aSorted.begin(), aSorted.end());
        sort(bSorted.begin(), bSorted.end());

        long long option2 = c;
        bool option2Possible = true;

        for (int i = 0; i < n; i++) {
            if (aSorted[i] < bSorted[i]) {
                option2Possible = false;
            } else {
                option2 = option2 + (aSorted[i] - bSorted[i]);
            }
        }

        // Pick the smaller valid option.
        long long best = -1;

        if (option1Possible) {
            best = option1;
        }

        if (option2Possible) {
            if (best == -1 || option2 < best) {
                best = option2;
            }
        }

        cout << best << "\n";
    }

    return 0;
}
