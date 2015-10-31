#Introduction to Cellular Communications
Speaker: Joe Blakenship
Date: September 12, 2015
Github: joeblakenship1

## Common Theme:
	-Most this can be broken by NSA

## History:
	-1920s: Detroit Police Dept utilizes a 2MHz system (literally a trunk of radio stuff)
	-1940s:
		-Frequencies were allocated between 30 & 40 MHz.
		-St. Louis deploys first public mobile phone (three channels in 150Mhz)
			-FCC allocated 6 channels, 60Khz apart (had massive interference)
		-Boston & NY operated 35 & 44 MHz in 1947 (originally Push To Talk)
	-1950s:
		-AT&T proposed a broadband mobile phone system operating in 800MHz range
		-Wire line channels at 150MHz expanded from 5 to 11
	-1970s:
		-AT&T started detailing the original cellular system
		-Chicago was the first "telephonized" city
	-1980s:
		-FCC issues standards

## Definitions:
	-Half Duplex: forward and backward take turns (very uncommon)
	-Full Duplex: forward and backwords simultaneously
	-Forward Channel: base station to mobile (down link)
	-Reverse Channel: Mobile to base (up link)
	-Cell Site/Base Station
		-Physical location provides coverage
		-Includes power sources, interface equipment, radio frequency transmitters/ receivers, antenna systems
		-Towers can be used by multiple carriers and in multiple formats
	-Registration: process of cell phones turning on and reporting to cell site
	-System ID (SID): unique identifier for each cellular network
	-Voice Comm:
		-Forward Voice Channel (FVC), Reverse Voice Channel (RVC)
		-Consists of voice audio traffic and data
			*Caller ID
	-Analog: Sine waves
	-Digital: Sine wave with clipping for beautiful highs and lows
	-MTSO
		-Base station -> MTSO
		-Manages scores of cells and base stations
		-Uses DBs to authenticate phones and customers
	-PSTN (Public Switched Telephone Network)
		-Pretty old tech but getting converted to fiber
		-Connects phones worldwide
		-Made up of local networks, exchange are networks and long-haul networks
		-Radio stuff to landline stuff
	-Cell site
		-Cellular tech (re)uses indivudal radio frequencies by dividing a service area into separate geographic zones called cells
		-Cells can be as small as a building up to 20+ miles across
		-Each cell is equiped with its own radio transmitter/receiver atenna array
		-Connects to MTSO on a basic level, more connections between them
	-Sectoring:
		-Splitting a cell into sectors by replacing an *omnidirectional* base station antenna w/ several directional antennas on the same tower
		-This divides a single cell coverage into 3, 6 or more distinct areas
			-120 degrees for 3 & 60 for 6 around the site
			-Each with their own frequences
		-Hexagonal shaped
	-Clustering
		-Must use diff freqs so there is no overlap of freqs, so handoff works a->b =/= a->
		-Cellular systems allocate all available frequences among a group of cells
	-Global System for Mobule Communications (GSM)
		-Most common world wide as of 2014
		-Most 2g GSM networks operating in 900 MHz or 1800 MHz band
		-Time slots
		-Used IS-136
		-Uses variety of vocoders (EFR, UMTS, AMR-Narrowband)
		-Uses SIM card

## Formats & Standards:
	-Frequency Re-use Plan
		-FMDA (practically deprecated): Very good for low density
		-TDMA
		-CDMA
	-Format/Standard
		-AMPS, IS-136, 4G
	-Operating Frequency
		-Various

## Network Components:
	-Frequency-Division Multiple Access
		-Available Spectrum is divided into small blocks of frequences
		-Each channel is assigned to one of these blocks
		-Designated block per call
	-Time-Division Multiple Access:
		-Alternators 4 different people through a channel, switches so quickly four people could use it
	-Code Division Multiple Access:
		-For entire spectrum "chopped" into multiple codes so people could use a whole spectrum
		-As higher usage user gets to use less code
		-Very efficient for maintaining
		-Uses soft hand offs, as you go from one cell to another cell weens you off the cells and onto the other's

## Standards:
	-Also called formats
	-Generally driven by NGO then adopted by Governments
	-Standards are agreed upon by protocols set by various organizations

## Stuff:
	-1st gen phones used analog FM to transmit voices
	-2nd gen phones used digital tech - Digital Base Band allows smart phone capabilities
	-Smart phones: Cell Network Layer -> Device Layer -> OS Layer -> UI/UX App Layer
		-The radio is just a peripheral that is connected through the micro-control unit
		-Square Card reader is read as audio (it goes in through headphone jack), injectible due to RAM being on baseband layer
		-Hardware vector is less commonly attacked
	-Hard hand offs almost never happen
	-Network: Mobile telephone switching office (MTSO), cell site, and format/standard (most important stuff)
	-Enhanced Data rates for GSM Evolution (EDGE). Considered pre-3G radio tech. Designed as 2G compatible.  (2.5G/3G tech)
	-America really fucked the world with GSM vs CMDA
	-4G is where a lot of the world used WiMAX2 (makes heavier use of WiFi tech) the USA decided to continue using Radio Waves for advanced stuff
	-A lot of the world is led by their military frequencies
	-There is FOSS to make pico towers using software
	-If phone is off but battery is still in the phone continues looking for control channel
	-Encryption is always lowest common denominator

## Network Functionality:
	-Registration
		-Registration most occur before doing anything else
		-On bootup does an IO check from Digital->Analog while Smartphone is booting up, eventually cross over occurs
		-Cellphone is always looking for signals (promiscuis)
		-If phone has no work, sits in "idle" mode
		-Identifies self on the RCC (reverse carrier channel)
		-Cell sites relay this information to MTSO
	-Supervisory Audio Tone (SAT)
		-High-pitched, inaudible tone that helps the system distinguish between callers on the same channel but in different cells
		-The SAT is transmitted continuously during a call (voice channel)
		-The SAT can be muted when a phone is powered up/down or site sends a "burst" of data
	-Mobile Identification Number (MIN)
		-aka mobile subscription identification number (MSIN)
		-phone number
	-International Mobile Station Equipment Identity (IMEI)
		-Number to identify 3GPP and iDEN mobile phones, as well as some sat phones
		-Type Allocation Code (TAC)
		-Final Assembly Code (FAC) (ceased after 2004)
	-International Mobile Subscriber Identity (IMSI)
		-Identify the user of a cellular network and is a unique identification associated with all cellular networks
		-64Bit field
		-To prevent eavesdropping identifying and tracking the subscribr on the radio interface the IMSI is rarely sent, instead uses TMSI
	-Phone constantly looking for new tower
	-Towers looking for new phones
	-If SAT is not there call will end (can be done by jamming)

## CallerID:
	-Sent to the called party by the telephone switch as an analog data stream between 1st and 2nd rings
	-If phone call is answered too quickly after first ring, caller ID info will not be transmitted
	-Two types: name and name+number
	-Can be spoofed

## FOSS:
	-Openmako: Project to create a family of open source mobile phones
	-OpenBTS: Software based GSM access point, allowing standard GSM compatible mobile phones to be used as SIP endpoits in VOIP networks
	-OsmocomBB: Free firmware for the devices known as "Baseband Processor" chipsets found in GSM mobile phones
	-PiPhone: Tutorial and material on adafruit
	-Software Defined Radio: Projects to use radio over software
	-AirProbe: Part of PyBomb package

## Resources
	-OpenSignal
	-WiGle
