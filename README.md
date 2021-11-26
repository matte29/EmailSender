# EmailSender

An application that can be used to send emails to many recipients in a short amount of time.

## What you need

- An HTML file for the body of the email

- A CSV file for the list of recipients (email addresses)


## Bugs

Doesn't work when it trys to read the values it fails

Error message: 2021/11/26 13:25:10 http: panic serving [::1]:65257: runtime error: invalid memory address or nil pointer dereference
## Future features

I don't know how to word this so here is an example.

In the body of the HTML file you can have a {Name} value that can be filled with values from the CSV file.

![Example](https://i.imgur.com/riTbrp0.png "Example")

It will take the value from the CSV file and use it to fill that value in HTML that is shown in the email.
