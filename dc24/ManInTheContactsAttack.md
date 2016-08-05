# Man In the Contact Attack

Day: August 05, 2016
By: @SecuringApps

# Main

## Introduction
- Popular messaging apps switch to e2e
	- great comm around it
	- privacy now is a requirement
- Debates at the government level to ask for backdoors
	- Going Dark?
	- Terorists?
- Increased feeling these apps are "unbreakable"

## Super Crypto but wait
- Advanced ratcheting in Signal Protocoal
- But how messaging apps auth myself?
	- No explicit identifier
	- provisioning done by SMS
	- Link to device
	- When I change phone #?
- And my contacts?
	- Get them from my address?
	- No manual import

## Threat Model: mobile focus
- Focus on the contacts
- Contacts accessible to all apps

## Accessing Contacts
- Easy to read/modify/create contacts
	- There is an API for that
	- Only challenge is "crappy documentation"
- Shared DS accessible in read /write
	- Only restricted by permissions
	- Contains auth data in clear!
- Room for side channel attack
	- Not requiring a rooted device

## Contact List Attacks
- Write an app to switch phone numbers of friends 
- Can also create contacts with whitespace b/c the apps remove the whitespace 
- Proxy via MITC app on phone

## Results
- WhatsApp: Notification received and does not change name to be correct
- Telegram: Same ""
- Signal: Same ""
- Creating << Foo>> is far more discrete than just switching users
	- Whitespace is not visible
	- Requires new contact but MITC attacks can delete/recreate << Foo>> as often as needed
- WhatsApp: Possible to share a real conversation between Bob and Eve via Eve
- Telegram: Same, s long as the new contacts are used for the first time
- Signal: Same results
	- Web version requires Android Phone to login
	- Phone number always displayed below contact name

## Risk Assessment
- Simple evaluation: risk = easiness of attack * user impact
- Difficulty: Low- Medium
	- Technically Low: Easy to access contacts, not problem to get MITC app approved
	- Logistics: Medium: One phone number is enough, need to convince users to download
- Impact: High: Thousands of users can get spied on

## Countermeasures:
- Messaging app
	- Display phone number on contact
	- Give up implicit trust of contracts
- Mobile OS
	- restrict for accessing contacts
- End User
	- Check your contact
	- Beware of new conversations
	- Avoid installing apps asking for modify permissions
- E2E can't guarantee privacy if not sure who you're talking to
- Security Model around contacts == way to big
- Auth the other party
