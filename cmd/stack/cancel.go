package main

// cancel flow:
// 0. parse optional arguments
// 1. get PAT from XDG dir
// 2. build POST payload
// 3. get all my active releases
// 4. for each release:
// 4.1 cancel
// 4.2 abandon
// 5. print IDs of cancelled/abandoned releases
