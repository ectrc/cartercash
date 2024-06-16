import buttons from "./frame/buttons.module.css";
import buttonColours from "./colours/buttons.module.css";

import games from "./browser/game.module.css";
import gameColours from "./colours/games.module.css";

const styles = {
  buttons: {
    collect: `${buttons.frameButtonOuter} ${buttonColours.collect}`,
    green: `${buttons.frameButtonOuter} ${buttonColours.green}`,
  },
  games: {
    classic: `${games.classicGameCard} ${gameColours.blue} ${gameColours.scanline}`,
    mega_diamond: `${games.megaDiamondGameCard} ${gameColours.purple} ${gameColours.sunrise}`,
    mexican_chilli: `${games.megaDiamondGameCard} ${gameColours.fire}`,
  },
};

export default styles;
