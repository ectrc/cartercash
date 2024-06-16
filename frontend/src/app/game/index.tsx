import { useNavigate } from "@tanstack/react-router";

import styles from "src/styles/games/index.module.css";
import { CollectButton } from "../frame/buttons";

// const games: Record<string, React.ReactNode> = {};

const Game = () => {
  const navigate = useNavigate();

  const handleReturn = () => {
    navigate({ to: "/" });
  };

  return (
    <div className={styles.gameContainer}>
      <button onClick={handleReturn} className={styles.goBackToBrowser}>
        <CollectButton text="Go Back" />
      </button>
    </div>
  );
};

export default Game;
