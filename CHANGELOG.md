# Changelog

## 0.2.0
- Remove `FPS Stability` chart
    - It shows misleading data, a high variance in FPS does not mean a high variance in perceived smoothness/performance in the same way it does with frametimes. Example: FPS constantly jumping around between 100-140 is "high variance" but actually only 2.9ms swings(10ms-7.1ms) and FPS jumping around between 50-70 would appear as lower variance but it isn't, that's 5.7ms swings (20ms-14.3ms) which is in fact higher variance and will feel much choppier. Variance/std dev in FPS is not relevant, only frametimes.
- Remove other misleading/contradictory charts
    - Only show max/avg/min bar chart for FPS, not frametime. People generally want to see how much "more FPS" they're getting, not how much "less time it takes to render a frame". Overall frametime average can still be found in the summary tab.
    - Only show relative average FPS % improvement chart for similar reasons as above.
    - Only show frametime line graph. FPS line graph can still be found in the all data tab.
    - Only show frametime density graph. These can kind of illustrate the same thing (a tighter/narrower distribution is better in both cases) but the FPS density graph is just harder to actually look at and decipher/draw conclusions from.
- Consolidate FPS & frametime tabs into one `Performance` tab, showing the more accurate/relevant metric for each chart.

## 0.1.2
- Fix docker builds

## 0.1.1
- Fix the bars of the relative average charts being disproportionate to the actual % change
- Change the labels of Min/Avg/Max charts to Max/Avg/Min to reflect the order of the bars in the chart

## 0.1.0
- Allow users to edit their uploaded benchmark details
- Sort FPS & frametime bar charts from worst to best
- Display max length of input fields in benchmark creation/edit forms
- Limit benchmark file names to 50 characters
- Adjusted mangohud configuration recommendations
- Add instructions to setup RTSS settings
  - Allows Linux & Windows line graphs to line up properly
- Fix relative average comparisons not always based on the actual lowest scoring baseline 
- Fix incorrect RAM humanization
