# go-mail-client
simple mail client support sending html content type and attachment

noauth example:  
itg-email-client -attach script1.log -content html -subject "test using go" -to heince@gmail.com -from heince -body /tmp/my.html -noauth

with auth example:
itg-email-client -attach export.txt -content html -subject "test auth" -to heince@gmail.com -from cs.blablaxx@gmail.com -body /tmp/my.html -host smtp.gmail.com -port 587 -user cs.blablaxx@gmail.com -password mypassword

