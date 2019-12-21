Download/Clone the go server from github.
Make sure you have go installed and configured.

In your terminal access the folder by going in the directory of the code.

The command 'go build RestAPI_library_server.go' will configure the server.
However in order to run it you must use the command './RestAPI_library_server'

The server is now on and you can go to a browser, preferablly chrome and type in 'localhost:7777/'

To search a book ID type 'localhost:7777/book/{bookID between 1 and 6}'
To find the most popular book type 'localhost:7777/popular'
To find the most demanded book type 'localhost:7777/mostIssued'
To find all available books type 'localhost:7777/availablebooks'
To find all issued books type 'localhost:7777/issuedbooks'

This takes you to the homepage.
To go to the other pages, type in 'status/','demanded/','popular/' or 'checkedOutUser/' after the URL

To stop the server (in mac OS X) press 'control' + 'C'
