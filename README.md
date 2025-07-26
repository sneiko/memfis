### PProf Heap Analyzer

A lightweight tool for visualizing Go heap profile reports from saved pprof files.
*Key Features*
- File-based analysis: Load pre-captured heap profiles without a running application
- Interactive dashboard: Visualize memory statistics with charts and tables
- Standalone: No need for live application connection - works directly with profile heap report files

Usage with --file Flag
bash
```
./pprof-analyzer --file <path-to-profile>
```
