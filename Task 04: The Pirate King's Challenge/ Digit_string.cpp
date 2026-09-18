#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t > 0) {
        t--;

        string s;
        cin >> s;
        int n = s.size();

        // count how many '1's and '3's are in the whole string.
        int onesThreesRemaining = 0;
        for (int i = 0; i < n; i++) {
            if (s[i] == '1' || s[i] == '3') {
                onesThreesRemaining++;
            }
        }

        int twosSoFar = 0;
        int best = twosSoFar + onesThreesRemaining;

        for (int i = 0; i < n; i++) {
            char c = s[i];

            if (c == '2') {
                twosSoFar++;
            } else if (c == '1' || c == '3') {
                onesThreesRemaining--;
            }

            int kept = twosSoFar + onesThreesRemaining;
            if (kept > best) {
                best = kept;
            }
        }

        int deletions = n - best;
        cout << deletions << "\n";
    }

    return 0;
}
