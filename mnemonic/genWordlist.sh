echo -e "package mnemonic\n\n//Words list of Mnemonics\nvar Words []string = []string{" > mnemonic.go
cat eff_large_wordlist.txt | sed -rn 's/^.*\t(.*)$/\t"\1",/p' >> mnemonic.go
echo "}" >> mnemonic.go
