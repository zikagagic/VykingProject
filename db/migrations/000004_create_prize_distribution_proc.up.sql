CREATE PROCEDURE distribute_prizes(IN tournamentID CHAR(36))
BEGIN
    DECLARE totalPrize DECIMAL(12,2);

    -- Get the total prize pool for the tournament
    SELECT prize_pool INTO totalPrize
    FROM tournament
    WHERE id = tournamentID;

    -- Distribute to first place (50%)
    UPDATE player
    JOIN prize_placement ON player.id = prize_placement.playerID
    SET player.account_balance = player.account_balance + (totalPrize * 0.5)
    WHERE prize_placement.tournamentID = tournamentID
      AND prize_placement.placement = 1;

    -- Distribute to second place (30%)
    UPDATE player
    JOIN prize_placement ON player.id = prize_placement.playerID
    SET player.account_balance = player.account_balance + (totalPrize * 0.3)
    WHERE prize_placement.tournamentID = tournamentID
      AND prize_placement.placement = 2;

    -- Distribute to third place (20%)
    UPDATE player
    JOIN prize_placement ON player.id = prize_placement.playerID
    SET player.account_balance = player.account_balance + (totalPrize * 0.2)
    WHERE prize_placement.tournamentID = tournamentID
      AND prize_placement.placement = 3;
END 