# Hints

## General

- [Use the methods found in the time package.][time]

## 1. Parse appointment date

- There is a [method][time.parse] for parsing a `string` into a `Time`.

## 2. Check if an appointment has already passed

- There are [methods][before] for [comparing][after] `Times` and [getting][now] the current date and time.

## 3. Check if appointment is in the afternoon

- There is a [method][hour] for getting the hour of a `Time`.

## 4. Describe the time and date of the appointment

- Convert the given string to a `Time` then format the answer string accordingly, using the appropriate [methods][time] to extract the needed constituents.

## 5. Return the anniversary date

- Create a `Time` of the anniversary date.

[time]: https://golang.org/pkg/time/#pkg-index
[time.parse]: https://golang.org/pkg/time/#Parse
[before]: https://golang.org/pkg/time/#Time.Before
[after]: https://golang.org/pkg/time/#Time.After
[now]: https://golang.org/pkg/time/#Now
[hour]: https://golang.org/pkg/time/#Time.Hour
