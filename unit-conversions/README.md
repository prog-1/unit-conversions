# Conversions

The repository contains a boilerplate code for programs that do some sort of a
unit conversion or mathematical calculations.

For each program there are tests being executed on the push and pull request to
the repository. The tests for a correct program must pass.

1. Use **only** `float64` variables in all programs.

1. Use the following functions:

   * [fmt.Println](https://pkg.go.dev/fmt#Println)
   * [fmt.Scan](https://pkg.go.dev/fmt#Scan)

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

Directory `rectangle`.

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

Directory `circle`.

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

Directory `speed`.

The program must convert the speed in `km/h` to the speed in `m/s`.

Example:

```
The program converts km/h to m/s.
Enter the speed in km/h:
100
The speed in m/s: 27.77777777777778
```

### Temperature (2 points)

Directory `temperature`.

The program must convert Celsius degrees to Fahrenheit degrees. Finding a
conversion formula is part of the exercise.

Example:

```
The program converts temperature from Celsius to Fahrenheit.
Enter the temperature in Celsius:
100
The temperature in Fahrenheit: 212
```

### Quadratic equation (4 points)

Directory `quadratic`.

Given coefficients `A`, `B`, and `C` of a quadratic equation `Ax^2 + Bx + C =
0`, write a program that prints the equation solutions `x1` and `x2`.

You will need to use [math.Sqrt](https://pkg.go.dev/math#Sqrt) function to calculate a square root.


Example:

```
The program finds solutions of a quadratic equation.
Enter the coefficients A, B and C:
1 3 -10
x1 = 2
x2 = -5
```

### Max and min (8 points)

Directory `maxmin`.

Write 4 programs (each in a separate subdirectory):

1. (`maxmin/max2`) a program that calculates the maximum of two numbers. You
   can use the following formula: `max(a, b) = (a + b + |a - b|)/2`. You will
   need to use [math.Abs](https://pkg.go.dev/math#Abs) function to calculate an
   absolute value. Example:

   ```
   The program finds the maximum of two numbers max(a, b).
   Enter two numbers:
   5 7
   max: 7
   ```

1. (`maxmin/min2`) a program that calculates the minimum of two numbers `min(a, b)`,

1. (`maxmin/max3`) a program that calculates the maximum of three numbers `max(a, b, c)`,

1. (`maxmin/min3`) a program that calculates the minimum of three numbers `min(a, b, c)`,

### Resistance (3 points)

Directory `resistance`.

Write a program that calculates the resistance `R` of two resistors with the
resistance `R1` and `R2` assuming that

1. the resistors `R1` and `R2` are connected in series, i.e. `R = R1 + R2`;
1. the resistors are connected in parallel, i.e. `1 / R = 1 / R1 + 1 / R2`.

Example:

```
The program calculates the resistance R of two resistors.
Enter the resistance of two resistors:
4 8
Resistance when connected in series is 12; resistance when connected in parallel is 2.6666666666666665.
```
