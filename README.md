# abobus

The repository contains a boilerplate code for programs that do some sort of a
unit conversion or mathematical calculations.

For each program there are tests being executed on the push and pull request to
the repository. The tests for a correct program must pass.

1. Use **only** `float64` variables in all programs.

1. Use the following functions:

   * [fmt.Println](https://pkg.go.dev/fmt#Println)
   * [fmt.Scan](https://pkg.go.dev/fmt#Scan)

1. While printing `float64` values, use formatting with 2 digits precision (see [example](https://play.golang.org/p/m7miCyzIaie)).

## Submission

The workflow to complete and submit the assignment is the following:

1. In the top-right corner of the page, click **Fork**

![fork button](https://docs.github.com/assets/images/help/repository/fork_button.jpg)

1. The repository will be "copied" to your GitHub account.

1. Complete the assignment by commiting your changes to your GitHub repository.
   Make sure that the tests in **Actions** tab of the repository pass successfully.

1. Once done, create a **Pull request** to the original [prog-1/unit-conversions](https://github.com/prog-1/unit-conversions) repository.

   * Select **Pull requests** tab in your repository.
   * Click **New pull request**.
   * Click **Create pull request**.

## Programs

### Rectangle (2 points)

The program must print the perimeter and the square of a rectangle given the
rectangle (non-opposite) sides.

Example:

```
The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the breadth of the rectangle:
4 5.5
Perimeter: 19
Square: 22
```

### Circle (2 points)

The programm must print the perimeter and the square of a circle given the
circle radius.

Example:

```
The program prints the perimeter and the square of a circle given the radius.
Enter the radius:
8
Perimeter: 50.26548245743669
Square: 201.06192982974676
```

You will need to use [math.Pi](https://pkg.go.dev/math#pkg-constants) constant to represent `π` value.

### Speed (2 points)

The program must convert the speed in `km/h` to the speed in `m/s`.

Example:

```
The program converts km/h to m/s.
Enter the speed in km/h:
100
The speed in m/s: 27.77777777777778
```

### Temperature (2 points)

The program must convert Celsius degrees to Fahrenheit degrees. Finding a
conversion formula is part of the exercise.

Example:

```
The program converts temperature from Celsius to Fahrenheit.
Enter the temperature in Celsius:
100
The temperature in Fahrenheit: 212
```
