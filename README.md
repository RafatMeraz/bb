# bb - BoroBhai
### CLI tool for generate cubit/bloc and other stuffs

## Install
To install ``bb`` in your mac, use these commands.

```
brew tap RafatMeraz/bb
brew install bb
```

## Generate

### cubit
To generate cubit use

``bb generate cubit NameOfCubit StateName``

Here, Cubit name if mandatory, but state name is optional. If you want to write a cubit with primitive type
then just add primitive tag in command, which is

``bb generate cubit NameOfCubit --p=String``

By default, if you left the state name empty and haven't added any primitive tag then it will add ``int`` as 
default data type for that cubit.

## bloc
To generate a bloc use

``bb generate bloc NameOfBloc --state=StateName --event=EventName``

or use shorthand tag

``bb generate bloc NameOfBloc -s StateName -e EventName``

# state
To generate only state use

``bb generate state NameOfState``

if it is part of any bloc/cubit just mention it by tag

``bb generate state NameOfState --partOf=sample_cubit.dart``

or in shorthand

``bb generate state NameOfState -p sample_cubit.dart``

# event
To generate only event use

``bb generate event NameOfEvent``

if it is part of any bloc/cubit just mention it by tag

``bb generate event NameOfEvent --partOf=sample_cubit.dart``

or in shorthand

``bb generate event NameOfEvent -p sample_cubit.dart``