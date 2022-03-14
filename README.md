# gender

A Go package for identifying gender.

Test it out at [gender.hankstoever.com](http://gender.hankstoever.com)

Data from [US Social Security](http://www.ssa.gov/oact/babynames/limits.html)

## Usage

### Classify

~~~bash
go run learn.go
~~~

Come back and test your results.

### Check

~~~bash
go run names/test.go
~~~

Output:

~~~
hank Male 0.9999999648880282
mark Male 0.668767623154752
hannah Female 0.9219143576826196
rachael Female 0.975310158457192
edward Male 0.7647226173541962
norah Female 0.9908132017693094
henry Male 0.8595182263909614
charlie Male 0.6081906629459236
ben Male 0.9850965062301491
claire Female 0.9584026622296173
matt Male 0.9990551181102363
lauren Female 0.709172456248351
marsha Female 0.9678191887441999
~~~

You can also check against Census data from 2000:

~~~bash
go run accuracy/accuracy.go
~~~

Output:

~~~
Rows: 29763
Wrong: 2434
Accuracy: 0.918220609481571
~~~

You can also test out your classifier with a web interface:

~~~bash
go run web/web.go
~~~

Then open [http://localhost:5000](http://localhost:5000) in your browser.

![screenshot](https://f.cloud.github.com/assets/1109058/2061415/0c78d27c-8c51-11e3-8ce3-7106912414ec.png)

## Analysis

`**** Data is only a sample and not representative of [UP Global](http://up.co)'s official records ****`

The purpose of this project was to get a better understanding of female participation at [startup weekend](http://startupweekend.org).

Every event has attendees, organizers, speakers, coaches, and judges. We have historical data for this, but `gender` was never something we stored. This is why we'll be classifying based on first name.

### Volunteers

`go run startup_weekend/startup_weekend.go`

Output:

~~~
Organizer
2011
Total: 45
Female: 21 ( 46.666666666666664 %)
Male: 24
2012
Total: 1555
Female: 659 ( 42.37942122186495 %)
Male: 896
2013
Total: 2968
Female: 1379 ( 46.4622641509434 %)
Male: 1589

Mentor
2011
Total: 146
Female: 44 ( 30.136986301369863 %)
Male: 102
2012
Total: 2980
Female: 944 ( 31.67785234899329 %)
Male: 2036
2013
Total: 4438
Female: 1475 ( 33.23569175304191 %)
Male: 2963

Speaker
2011
Total: 32
Female: 7 ( 21.875 %)
Male: 25
2012
Total: 457
Female: 150 ( 32.822757111597376 %)
Male: 307
2013
Total: 796
Female: 297 ( 37.31155778894472 %)
Male: 499

Judge
2011
Total: 42
Female: 11 ( 26.190476190476193 %)
Male: 31
2012
Total: 1304
Female: 442 ( 33.895705521472394 %)
Male: 862
2013
Total: 1970
Female: 685 ( 34.77157360406091 %)
Male: 1285
~~~

### Attendees

`go run attendees/attendees.go`

Output:

~~~
2009
Total: 1034
Female: 246 ( 23.791102514506772 %)
Male: 788

2010
Total: 7467
Female: 2258 ( 30.2397214410071 %)
Male: 5209

2011
Total: 21083
Female: 6756 ( 32.044775411468954 %)
Male: 14327

2012
Total: 57349
Female: 19896 ( 34.69284555964359 %)
Male: 37453

2013
Total: 56789
Female: 20493 ( 36.086213879448486 %)
Male: 36296

2014
Total: 2904
Female: 1185 ( 40.80578512396694 %)
Male: 1719
~~~

