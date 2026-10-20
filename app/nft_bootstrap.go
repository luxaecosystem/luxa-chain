package app

import (
	"fmt"
	"cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GenesisNFTClassID  = "luxa-relics"
	GenesisNFTTreasury = "luxa1eg6d8axpw2t3en2g8t0g5qtj4wm2fh3q4tmkue"
)

func (app *App) BootstrapGenesisNFTs(ctx sdk.Context) error {
	if app.NFTKeeper.HasClass(ctx, GenesisNFTClassID) {
		return nil
	}

	app.Logger().Info(">>> [GENESIS BOOTSTRAP] Registrazione luxa-relics e conio 5 NFT... <<<")

	class := nft.Class{
		Id:          GenesisNFTClassID,
		Name:        "LUXA Genesis Relics",
		Symbol:      "LXR",
		Description: "Official Genesis Artifacts & Citizen Zero of the LUXA Sovereign Chain",
		Uri:         "https://luxaecosystem.alwaysdata.net/nftlist.json",
	}

	if err := app.NFTKeeper.SaveClass(ctx, class); err != nil {
		fmt.Printf("Errore SaveClass: %v\n", err)
		return err
	}

	treasuryAddr, err := sdk.AccAddressFromBech32(GenesisNFTTreasury)
	if err != nil {
		fmt.Printf("Indirizzo treasury non valido: %v\n", err)
		return err
	}

	relics := []nft.NFT{
		{
			ClassId: GenesisNFTClassID,
			Id:      "0",
			Uri:     "https://luxaecosystem.alwaysdata.net/assets/nfts/citizen_zero.jpeg",
		},
		{
			ClassId: GenesisNFTClassID,
			Id:      "1",
			Uri:     "https://luxaecosystem.alwaysdata.net/assets/nfts/chrono_key_master.jpeg",
		},
		{
			ClassId: GenesisNFTClassID,
			Id:      "2",
			Uri:     "https://luxaecosystem.alwaysdata.net/assets/nfts/cyber_shadow_node.jpeg",
		},
		{
			ClassId: GenesisNFTClassID,
			Id:      "3",
			Uri:     "https://luxaecosystem.alwaysdata.net/assets/nfts/grandmaster_of_servers.jpeg",
		},
		{
			ClassId: GenesisNFTClassID,
			Id:      "4",
			Uri:     "https://luxaecosystem.alwaysdata.net/assets/nfts/neon_data_valkyrie.jpeg",
		},
	}

	for _, relic := range relics {
		if err := app.NFTKeeper.Mint(ctx, relic, treasuryAddr); err != nil {
			fmt.Printf("Errore Mint relic %s: %v\n", relic.Id, err)
			return err
		}
	}

	fmt.Println(">>> [GENESIS BOOTSTRAP] Successo: 5 NFT coniati per luxatreasury! <<<")
	app.Logger().Info(">>> [GENESIS BOOTSTRAP] Successo: 5 NFT coniati per luxatreasury! <<<")
	return nil
}
