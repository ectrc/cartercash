import { useNavigate } from "@tanstack/react-router";

import styles from "src/styles";

type BaseCardProps = {
  id: string;
} & React.HTMLAttributes<HTMLDivElement>;

const BaseCard = (props: BaseCardProps) => {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate({
      to: `/game/$id`,
      params: { id: props.id },
    });
  };

  return <div onClick={handleClick} {...props} />;
};

const ClassicGameCard = () => {
  return (
    <BaseCard id="classic" className={styles.games.classic}>
      Classic 777
    </BaseCard>
  );
};

const MegaDiamondGameCard = () => {
  return (
    <BaseCard id="diamonds" className={styles.games.mega_diamond}>
      Mega Diamond
    </BaseCard>
  );
};

const MexicanChilliGameCard = () => {
  return (
    <BaseCard id="mxchilli" className={styles.games.mexican_chilli}>
      Mexican Chilli
    </BaseCard>
  );
};

export { ClassicGameCard, MegaDiamondGameCard, MexicanChilliGameCard };
