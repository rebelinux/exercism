using System;

class RemoteControlCar
{
    private int rDistance = 0;
    private int usedBatt = 100;
    public static RemoteControlCar Buy()
    {
        return new RemoteControlCar();
    }

    public string DistanceDisplay()
    {
        return $"Driven {rDistance} meters";
    }

    public string BatteryDisplay()
    {
        if (usedBatt == 0) {
            return "Battery empty";
        } else { return $"Battery at {usedBatt}%"; }
    }

    public void Drive()
    {
        if (usedBatt > 0)
        {
            rDistance += 20;
        }
        if (usedBatt >= 1)
        {
            usedBatt--;
        }
       
    }
}
