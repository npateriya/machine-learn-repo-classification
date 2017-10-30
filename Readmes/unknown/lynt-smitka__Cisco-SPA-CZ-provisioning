# Cisco SPA CZ provisioning
Skript pro generování konfigurací pro telefony Cisco SPA

Vygeneruje konfigurační soubory pro zařízení uvedená v csv seznamu čísel (vzor v example.csv) podle šablony (vzor default.xml).

Vzorová šablona je připravena pro české prostředí - kodeky, dialplán, vyzvánění kódy, časy, lokalizace

Při provisioningu proběhne upgrade na aktuální firmware 7.6.2SR3. Jsou využívány služby http://tools.lynt.cz/provisioning.php.

Je vygenerován i telefonní seznam se všemi čísly - directory.xml.

Specifikace seznamu čísel:

`sériové číslo přístroje; linka; heslo; jméno; šablona konfigurace pro použití`

Použití:

`python generate.py ip_adresa_serveru csv_seznam_cisel`
