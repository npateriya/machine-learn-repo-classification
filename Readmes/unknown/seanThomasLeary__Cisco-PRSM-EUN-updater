# Cisco-PRSM-EUN-updater
Workaround for CSCut17139 so users can update the EUN message on CX sensors

This bug prevents the user from updating the End User Notification pages on Cisco CX sensors. 
The HTML which is normally embedded in the detail text becomes visible and does not perform its markup function.
The text gets saved to the config database and causes display problems in the actual EUN 
pages that are rendered to end users. Normally this type of problem could be addressed by a 
few Curl messages. However in this case the  EUN update also includes the caption text, EUN type, 
image binary, and other fields, since the fields of the database record are updated all 
together as a blob.

To allow users to work around the problem, the entire db record has to be retrieved, and then
sent back in an update, with only the text changed. This will at least fix the problem where 
users have inadvertently modified the EUN detail text. Future revisions will allow users to
modify other fields in the db record.

