class Lasagna
{
    public int ExpectedMinutesInOven() {
        return 40;
    }

    public int RemainingMinutesInOven(int rminutes) {
        return 40 - rminutes;
    }

    public int PreparationTimeInMinutes(int eminutes) {
        return 2 * eminutes;
    }

    public int ElapsedTimeInMinutes(int layers, int ominutes) {
        return PreparationTimeInMinutes(layers) + ominutes;
    }
}
